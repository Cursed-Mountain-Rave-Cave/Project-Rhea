import socket
import json
 
TCP_IP = '127.0.0.1'
TCP_PORT = 25565
BUFFER_SIZE = 1024
MESSAGE = ""
 
s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect((TCP_IP, TCP_PORT))

def send(login, password):
    loginJSON = json.dumps({"Login":login, "Password":password})
    msgJSON = json.dumps({"MessageType":"login", "Message": loginJSON})
    n = s.send(msgJSON.encode('utf-8'))

send("123", "123")
send("pl123", "pl123")

s.close()