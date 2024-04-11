import json
from random import randrange
from locust import HttpUser, task, between

class readFile():

    def __init__(self):
        self.data = []
    
    def getElement(self):
        size = len(self.data)
        if size > 0:
            index = randrange(0, size -1) if size > 1 else 0
            return self.data.pop(index)
        else:
            print("No hay elementos en la lista")
            return None

    def getFile(self):
        try:
            with open('traffic.json', 'r', encoding='utf-8') as file:
                self.data = json.loads(file.read())
        except Exception as e:
            print(e)

class WebsiteUser(HttpUser):
    wait_time = between(0.1,0.9)
    readFile = readFile()
    readFile.getFile()

    @task
    def sendRequest(self):
        data = self.readFile.getElement()
        if data is not None:
            res = self.client.post("/insert", json=data)
            response = res.json()
            print(response)
        else:
            print("No hay elementos en la lista")
            self.stop(True)