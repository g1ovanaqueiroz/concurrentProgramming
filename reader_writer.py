from Semaphore import *

empty = Semaphore(1)
reader_counter_flag = 0
mutex = Semaphore(1)

#writing
empty.wait()
#update_inode(inode)
empty.signal()

#reading
mutex.wait()
reader_counter_flag += 1
if (reader_counter_flag == 1):
    empty.wait()
mutex.signal()

#inode = get_inode()

#light switch
mutex.wait()
reader_counter_flag -= 1
if (reader_counter_flag == 0):
    empty.signal()
mutex.signal()