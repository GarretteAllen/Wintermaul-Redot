extends Node

signal login_success(data)
signal login_error(message)
signal register_success
signal register_error(message)

func _ready():
	WebSocketManager.connect("data_received", Callable(self, "_on_data_received"))
	
func _on_data_received(data):
	if data:
		if data.has("status"):
			match data["status"]:
				"logged_in":
					emit_signal("login_success", data)
				"error":
					if data.has("message"):
						if data["message"] == "Invalid password" or data["message"] == "User not found":
							emit_signal("login_error", data["message"])
						elif data["message"] == "Registration failed":
							emit_signal("register_error", data["message"])
				"registered":
					emit_signal("register_success")
		else:
			print("Status key not found")
	else:
		print("data is null")
