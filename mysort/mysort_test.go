package mysort



func TestQuickSort (t *testing.T) {
  dat := int32{3,5,6,3,2,1}
  result := int32{1,2,3,3,5,6}
  QuickSort(dat, 0, len(dat))
  if dat != result {
    t.Fatal("quick sort error")
  }
}
