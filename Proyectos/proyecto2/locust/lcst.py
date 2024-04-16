import json
from random import randrange
from locust import HttpUser, task, between

class readFile():

    def __init__(self):
        self.data = []

    def getElement(self):
        size = len(self.data)
        if size > 0:
            i = randrange(0, size-1) if size > 1 else 0
            return self.data.pop(i)
        else:
            print("No more elements to get")
            return None
    
    def getFile(self):
        try:
            with open('traffic.json', 'r', encoding='utf-8') as file:
                self.data = json.loads(file.read())
        except Exception as e:
            print(f"Error reading file: {e}")

class WebsiteUser(HttpUser):
    wait_time = between(0.1, 0.9)
    file = readFile()
    file.getFile()

    @task
    def sendRequest(self):
        element = self.file.getElement()
        if element is not None:
            res = self.client.post("/vote", json=element)
            response = res.json()
            print(response)
        else:
            print("No more elements to send")
            self.stop(True)