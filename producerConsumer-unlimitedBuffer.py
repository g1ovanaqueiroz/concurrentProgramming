from collections import deque

class Semaphore:
     
     def __init__(self, n):
         self.n = n

     def signal(self):
         self.n += 1
    
     def wait(self):
         self.n -= 1

mutex = Semaphore(1)
dataAvailable = Semaphore(0)

buffer = deque()

class Producer:

    data = "produceData()" #produceData()
    mutex.wait()
    buffer.add(data)
    dataAvailable.signal()
    mutex.signal()


class Consumer:

    dataAvailable.wait()
    mutex.wait()
    data = buffer.get()
    mutex.signal()
    #process(data)