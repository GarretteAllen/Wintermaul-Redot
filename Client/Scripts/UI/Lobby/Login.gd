extends Control

@onready var username_input = $Background/Username
@onready var password_input = $Background/Password
@onready var login_button = $Background/LoginButton
@onready var register_button = $Background/RegisterButton
@onready var lobby_scene = preload("res://Client/Scenes/UI/Lobby/LobbyBrowser.tscn")
@onready var ws_connected = WebSocketManager.ws_connected
var delay_timer = Timer.new()

func _ready():
	add_child(delay_timer)
	delay_timer.one_shot = true
	MessageHandler.connect("login_success", Callable(self, "_on_login_success"))
	MessageHandler.connect("login_error", Callable(self, "_on_login_error"))
	MessageHandler.connect("register_success", Callable(self, "_on_register_success"))
	MessageHandler.connect("register_error", Callable(self, "_on_register_error"))
	login_button.connect("pressed", Callable(self, "_on_login_pressed"))
	register_button.connect("pressed", Callable(self, "_on_register_pressed"))
	
func _on_login_pressed():
	var username = username_input.text.strip_edges()
	var password = password_input.text.strip_edges()

	if not WebSocketManager.ws_connected:
		WebSocketManager.connect_to_server()
		delay_timer.start(1.5)
		delay_timer.connect("timeout", Callable(self, "_send_login_message").bind(username, password))
	else:
		WebSocketManager.send_message("login", {"username": username, "password": password})


func _on_register_pressed():
	var username = username_input.text.strip_edges()
	var password = password_input.text.strip_edges()

	if not WebSocketManager.ws_connected:
		WebSocketManager.connect_to_server()
		delay_timer.start(1.5)
		delay_timer.connect("timeout", Callable(self, "_send_register_message").bind(username, password))
	else:
		WebSocketManager.send_message("register", {"username": username, "password": password})

func _on_login_success(data):
	print("Login successful! Redirecting to lobby browser...")
	get_tree().change_scene_to_packed(lobby_scene)

func _on_login_error(message):
	print("Login failed:", message)

func _on_register_success():
	print("Registration successful!")

func _on_register_error(message):
	print("Registration failed:", message)
	
func _send_login_message(username, password):
	WebSocketManager.send_message("login", {"username": username, "password": password})
	
func _send_register_message(username, password):
	WebSocketManager.send_message("register", {"username": username, "password": password})
