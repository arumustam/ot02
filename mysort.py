def quick_sort(dat, left, right, asc=True):
  def asc_helper(x,y): return x < y
  def des_helper(x,y): return x > y

  if right-left <= 1: return

  init_pivot_i = (left+right)//2
  last_i = right-1
  dat[init_pivot_i], dat[last_i] = dat[last_i], dat[init_pivot_i]

  dase_swap = des_helper
  if asc:
    dose_swap = asc_helper

  i = left
  for j in range(left, last_i):
    if dose_swap(dat[j], dat[last_i]):
      dat[i], dat[j] = dat[j], dat[i]
      i += 1
  dat[i], dat[last_i] = dat[last_i], dat[i]

  quick_sort(dat, left, i, asc)
  quick_sort(dat, i+1, right, asc)

if __name__ == '__main__':
  a = [3,1,2,5,7,4,6,7,9,4,8,9,5,5,3,8]
  print(a)
  quick_sort(a, 0, len(a))
  print(a)
