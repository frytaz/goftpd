-- check we have params
if not params then
	session:Reply(501, "Syntax: site addip <user> <mask> <...mask>")
	return false
end

-- site addip <user> <mask> <...mask>
if #params < 2 then
	session:Reply(501, "Syntax: site addip <user> <mask> <...mask>")
	return false
end

-- get current user, dont check for nil as it will error anyway
local caller = session:User()

-- get target user
local target, err = session:Auth():GetUser(params[1])
if err then
	session:Reply(500, "Error: " .. err:Error())
	return false
end

-- check permissions
if not acl:MatchTarget(caller, target) then
	session:Reply(500, "Permission denied")
	return false
end

-- update the user adding each mask
local err = session:Auth():UpdateUser(target.Name, function(u)
	for i, mask in params() do
		if i > 1 then
			local err = u:AddIP(mask)
			if err == nil then
				session:Reply(226, "Added Mask: " .. mask)
			else 
				session:Reply(501, "Unable to add mask '" .. mask .. "': " .. err:Error())
			end
		end
	end

	return nil
end)

if err ~= nil then
	session:Reply(500, "Error: " .. err:Error())
	return false
end

return true
