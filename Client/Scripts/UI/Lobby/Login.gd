extends Control

@onready var username_input = $Background/Username
@onready var password_input = $Background/Password
@onready var login_button = $Background/LoginButton
@onready var register_button = $Background/RegisterButton
@onready var lobby_scene = preload("res://Client/Scenes/UI/Lobby/LobbyBrowser.tscn")

func _ready():
	WebSocketManager.connect("data_received", Callable(self, "_on_data_received"))
	login_button.connect("pressed", Callable(self, "_on_login_pressed"))
	register_button.connect("pressed", Callable(self, "_on_register_pressed"))
	
func _on_login_pressed():
	var username = username_input.text.strip_edges()
	var password = password_input.text.strip_edges()
	
	if not WebSocketManager.is_connected:
		WebSocketManager.connect_to_server()
	WebSocketManager.send_message("login", {"username": username, "password": password})
	
func _on_register_pressed():
	var username = username_input.text.strip_edges()
	var password = password_input.text.strip_edges()
	
	if not WebSocketManager.is_connected:
		WebSocketManager.connect_to_server()
	WebSocketManager.send_message("register", {"username": username, "password": password})
	
func _on_data_received(data):
	if data["action"] == "login":
		if data["status"] == "success":
			print("Login successful! Redirecting to lobby browser...")
			get_tree().change_scene_to_packed(lobby_scene)
		else:
			print("Login failed:", data["message"])
	elif data["action"] == "register":
		if data["status"] == "success":
			print("Registration successful!")
		else:
			print("Registration failed:", data["message"])
