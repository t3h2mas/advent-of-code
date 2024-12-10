input = '2333133121414131402'

input = '12345'

input = '0110'

input = File.read('./09_input.txt').strip

AdventFile = Data.define(:id, :size, :free)

files = input.chars.each_slice(2).each_with_index.to_a.map do |(size, free), id|
  AdventFile.new(id:, size: size.to_i, free: free.to_i)
end

files

storage = []

files.each do |file|
  file.size.times do
    storage << file.id
  end
  file.free.times do 
    storage << '.'
  end
end;

res = storage.dup;

lo = 0
hi = res.length - 1

while lo <= hi
  # walk up to next free space
  while res[lo] != '.'
    lo += 1
  end

  #walk down to next file id
  while res[hi] == '.'
    hi -= 1
  end

  res[lo] = res[hi]
  res[hi] = '.'
end

res

res[lo], res[hi] = res[hi], res[lo]

checksum = 0

res.each.with_index { |id, i| break if id == '.'; checksum += id.to_i * i };

checksum.to_i
