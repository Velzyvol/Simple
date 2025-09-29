return {
	input = {
		[hash("p1")] = { left = hash("key_a"), right = hash("key_d"), down = hash("key_s"), up = hash("key_w") },
		[hash("p2")] = { left = hash("key_left"), right =  hash("key_right"), down = hash("key_down"), up = hash("key_up") }
	},
	position = {
		[hash("p1")] = vmath.vector3(),
		[hash("p2")] = vmath.vector3()
	}
}