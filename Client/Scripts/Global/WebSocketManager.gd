extends Node

var websocket = null
var ws_connected := false
const SERVER_URL = "ws://localhost:8080/ws"
var socket = WebSocketPeer.new()

signal connected()
signal disconnected()
signal data_received(data)

func connect_to_server():
	
	var err = socket.connect_to_url(SERVER_URL)
	if err != OK:
		print("Failed to connect to server:", err)
		return
	ws_connected = true
	print("Connecting to server..")
	emit_signal("connected")

func disconnect_from_server():
	if socket and socket.State == WebSocketPeer.STATE_OPEN:
		socket.close()
		print("Disconnected from server")
		emit_signal("disconnected")

func send_message(action, payload):
	if socket and socket.get_ready_state() == WebSocketPeer.STATE_OPEN:
		var message = {"action": action, "payload": payload}
		socket.send_text(JSON.stringify(message))
		print("Sent message:", message)
		
func _process(_delta):
	if ws_connected:
		socket.poll()
		while socket.get_available_packet_count() > 0:
			var packet = socket.get_packet()
			print("packet recieved")
			var packet_string = packet.get_string_from_utf8()
			var parsed = JSON.parse_string(packet_string)
			if parsed.has("status"):
				emit_signal("data_received", parsed)
			else:
				print("Received invalid JSON data")
