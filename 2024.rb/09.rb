input = '2333133121414131402'

input = '2833133121414131402'

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

# part one

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

# part two
files

# split files into files and free space
res = files
  .flat_map { |f| [AdventFile.new(id: f.id, size: f.size, free: 0), AdventFile.new(id: :free_space, size: 0, free: f.free)] }
  .filter { |f| !(f.id == :free_space && f.free == 0) }

res = input.chars.each_slice(2).each_with_index.flat_map do |(size, free), id|
  [id.to_s * size.to_i, ['.'] * free.to_i]
end.flatten

hi = res.length - 1

while hi >= 0
  if hi == res.length - 1
    p res
  end

  right_size = res[hi].size

  if res[hi] == '.'
    hi -= 1
    next
  end

  puts "hi: #{res[hi]}"

  lo = 0
  while lo < hi
    if res[lo] == '.'
      # walk right right_size spaces  
      middle = lo + 1

      while res[middle] == '.' && middle < hi && !(middle - lo == right_size)
        middle += 1
      end

      break if middle == hi
      break if middle - lo != right_size

      res[lo..middle], res[hi] = res[hi], res[lo..middle]
      p res
    end
    lo += 1
  end

  hi -= 1
end


checksum = 0

res.each.with_index { |id, i| next if id.start_with?('.'); checksum += id.to_i * i };

checksum.to_i
checksum = 0

res.each.with_index { |id, i| break if id == '.'; checksum += id.to_i * i };

checksum.to_i
