import socket
import json
import threading
import sys
from tkinter import *
from utils import *

class Application(Frame):
    def __init__(self, master=None):
        super().__init__(master)
        self.master = master
        self.pack()
        self.create_widgets()   
        self.configure(bg='#4c528f')
        
        self.server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        self.PORT = 3000
        self.IP_ADR = 'localhost'

        self.server.connect((self.IP_ADR, self.PORT))
        
        self.disconnected = False

        self.receive_thread = threading.Thread(target=receive, kwargs={'server':self.server, 'text':self.text})
        self.receive_thread.start()


    def create_widgets(self):
        self.text = Text(self, font='Arial 12', height=25, width=50, borderwidth=3, relief="groove")
        self.text.pack(side="top")

        self.text.tag_config('receive_all', foreground='black')
        self.text.tag_config('info', foreground='blue')
        self.text.tag_config('error', foreground='red')

        self.message = StringVar()
        self.entry = Text(self, font='Arial 12', height=2, width=50, borderwidth=3, relief="groove")
        self.entry.pack(side="top")

        self.button = Button(self)
        self.button["text"] = "Send"
        self.button["command"] = self.send
        self.button.pack(side="top")

        self.master.bind('<Return>', self.send)

        self.quit = Button(self, text="QUIT", fg="red",
                              command=self.master.destroy)
    

    def send_please(self, type, data):
        request = {'type':type, 'data':json.dumps(data)}
        self.server.send(bytes(json.dumps(request), 'utf-8'))                              

    def send(self, event=None):
        command = self.entry.get(1.0, END).split()

        if len(command) == 0:
            return

        if command[0] == '/register':
            if len(command) != 3:
                print("Incorrect register format")
            else:
                data = {'login':command[1], 'password':command[2]}
                self.send_please('register', data)

        elif command[0] == '/login':
            if len(command) != 3:
                print("Incorrect login format")
            else:
                data = {'login':command[1], 'password':command[2]}
                self.send_please('login', data)

        elif command[0] == '/exit':
            disconnected = True
            self.server.close()  
            self.receive_thread.join()                
            sys.exit() 

        elif command[0][0] != '\\':
            data = {'message':' '.join(command[0:])}
            self.send_please('send_all', data)  

        self.entry.delete(1.0, END) 


def receive(server, text):
    while True:
        try:
            msg = json.loads(server.recv(1024).decode('utf'))
            text.configure(state=NORMAL)
            if msg['type'] == 'receive_all':
                message = json.loads(msg['data'])

                text.insert(END, message['login'] + ': ' + message['message'] + '\n', 'receive_all')

            elif msg['type'] == 'error':
                message = json.loads(msg['data'])
                text.insert(END, message['info'] + '\n', 'error')

            elif msg['type'] == 'info':
                message = json.loads(msg['data'])
                text.insert(END, message['info'] + '\n', 'info')

            text.configure(state=DISABLED)
        except OSError:
            break

root = Tk()
root.title('Super Omega chat 2.0')
app = Application(master=root)
root.resizable(width=False, height=False)

app.mainloop()