import socket
import json
import utils
 
TCP_IP = '127.0.0.1'
TCP_PORT = 25565
BUFFER_SIZE = 1024
MESSAGE = ""
 
s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect((TCP_IP, TCP_PORT))

def login(login, password):
    loginJSON = json.dumps({"login":login, "password":password})
    msgJSON = json.dumps({"type":"login", "data": loginJSON})
    s.send(msgJSON.encode('utf-8'))

def register(login, password):
    loginJSON = json.dumps({"register":login, "password":password})
    msgJSON = json.dumps({"type":"register", "data": loginJSON})
    s.send(msgJSON.encode('utf-8'))

def receive():
    try:
        netData = utils.receiveJSON(s)
        
        print(''.join(netData))
        msg = json.loads(''.join(netData))
        print(msg)
        login = json.loads(msg['data'])
        print(login)
    except :
        print("Error!")


register("alex", "123")
receive()
login("alex", "123")
receive()

login("alex", "123")
receive()

register("alex2", "123")
receive()
login("alex2", "123")
receive()

s.close()