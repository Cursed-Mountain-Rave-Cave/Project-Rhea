import socket
import json
import utils
 
TCP_IP = '127.0.0.1'
TCP_PORT = 25565
BUFFER_SIZE = 1024
MESSAGE = ""
 
s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect((TCP_IP, TCP_PORT))

def send(login, password):
    loginJSON = json.dumps({"login":login, "password":password})
    msgJSON = json.dumps({"type":"login", "data": loginJSON})
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

send("12}3", "123")
print()
receive()
print('')
send("pl1{23", "pl123")
print('')
receive()
print('')
send("pl123", "pl123")
print('')
receive()
print('')

s.close()