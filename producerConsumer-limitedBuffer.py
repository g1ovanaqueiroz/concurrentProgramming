class Semaphore:
     
     def __init__(self, n):
         self.n = n

     def signal(self):
         self.n += 1
    
     def wait(self):
         self.n -= 1

mutex = Semaphore(1)
buffer = []
#bufferLimitAvailable = Semaphore(n)
dataAvailable = Semaphore(0)

#produce while the buffer isn't full, the buffer shall not reach maximum capacity 
class Producer:

    data = "produceData()" #produceData()
    #bufferLimitavailable.wait()
    mutex.wait()
    buffer.append(data)
    mutex.signal()
    dataAvailable.signal()

# consumer works only when there's produced data
class Consumer:

    dataAvailable.wait()
    mutex.wait()
    data = buffer.pop()
    mutex.signal()
    #bufferLimitAvailable.signal()
    #process(data)