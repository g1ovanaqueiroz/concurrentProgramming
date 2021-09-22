class Semaphore:
     
     def __init__(self, n):
         self.n = n

     def signal(self):
         self.n += 1
    
     def wait(self):
         self.n -= 1


class Barrier:
    
     counter = 0
     mutex = Semaphore(1)
     sema = Semaphore(0)

     def __init__(self, limit):
         self.limit = limit
    
     def wait(self):
         self.mutex.wait()
         self.counter += 1
         self.mutex.signal()

         if(self.counter < self.limit):
             self.sema.wait()
         self.sema.signal()