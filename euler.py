from time import time

def nums(v):
  res = 0
  for i in range(v):
    if i % 3 == 0 or i % 5 == 0:
      res += i
  return res
  
def fib():
  topval = 4000000 #4 million
  res = 0
  a = 1
  b = 0
  temp = 0
  while a < topval:
    temp = a
    a += b
    b = temp
    if a % 2 == 0:
      res += a
  return res
  
def main():
  start_t = time()
  
  result = nums(1000)
  print "Multiples of 3 and 5: " + str(result)
  result = fib()
  print "Sum of even fibs: " + str(result)
  end_t = time() - start_t
  
  print str(end_t*1000000) + "us"
  
main()
