extends Node

var websocket = null
var is_connected := false
const SERVER_URL = "ws://localhost:8081/ws"
var socket = WebSocketPeer.new()

signal connected()
signal disconnected()
signal data_received(data)

func connect_to_server():
	
	var err = websocket.connect_to_url(SERVER_URL)
	if err != OK:
		print("Failed to connect to server:", err)
		return
	is_connected = true
	print("Connecting to server..")
	emit_signal("connected")

func disconnect_from_server():
	if socket and socket.State == WebSocketPeer.STATE_OPEN:
		socket.close()
		print("Disconnected from server")
		emit_signal("disconnected")

func send_message(action, payload):
	if socket and socket.State == WebSocketPeer.STATE_OPEN:
		var message = {"action": action, "payload": payload}
		socket.send_text(JSON.stringify(message))
		print("Sent message:", message)
		
func _process(_delta):
	if is_connected:
		socket.poll()
		while socket.get_available_packet_count() > 0:
			var packet = socket.get_packet()
			var parsed = JSON.parse_string(packet)
			if parsed.error == OK:
				emit_signal("data_received", parsed.result)
			else:
				print("Received invalid JSON data")
