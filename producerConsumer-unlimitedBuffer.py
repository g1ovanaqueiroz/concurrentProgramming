from collections import deque
from Semaphore import *

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