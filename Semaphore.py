# you can pretend that's a real semaphore in python :p
class Semaphore:

    def __init__(self, n):
         self.n = n

    def signal(self):
         self.n += 1
    
    def wait(self):
         self.n -= 1