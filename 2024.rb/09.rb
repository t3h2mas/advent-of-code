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

hi = res.length - 1

while hi >= 0
  if hi == res.length - 1
    puts as_str(res)
  end

  right_size = res[hi].size

  p res[hi]

  if right_size == 0 || res[hi].id == :free_space
    hi -= 1
    next
  end

  lo = 0
  while lo < hi
    if res[lo].id == :free_space && res[lo].free >= right_size
      remaining_free_space = res[lo].free - right_size

      res[lo] = res[hi]
      res[hi] =  AdventFile.new(id: :free_space, size: 0, free: right_size)

      if remaining_free_space.positive?
        if res[lo + 1]&.id == :free_space
          res[lo + 1] = AdventFile.new(id: :free_space, size: 0, free: res[lo + 1].free + remaining_free_space)
        else
          res.insert(lo + 1, AdventFile.new(id: :free_space, size: 0, free: remaining_free_space))
        end
      end

      p as_str(res)
      break
    end
    lo += 1
  end

  hi -= 1
end

def as_str(s)
  s.map do |f|
    if f.id == :free_space
      '.' * f.free
    else
      f.id.to_s * f.size
    end
  end.join('')
end

def as_a(s)
  s.flat_map do |f|
    if f.id == :free_space
      ['.'] * f.free
    else
      [f.id.to_s] * f.size
    end
  end
end

as_str(res) == "00992111777.44.333....5555.6666.....8888.."


checksum = 0

as_a(res).each.with_index { |id, i| next if id.start_with?('.'); checksum += id.to_i * i };

checksum.to_i

as_a(res)[-1]

as_str(res)


class List
  class Node
    attr_reader :value
    attr_accessor :prev

    def initialize(value)
      @value = value
      @prev = nil
      @nnext = nil # keyword
    end

    def next=(val)
      @nnext = val
    end

    def next
      @nnext
    end

    def inspect
      "<Node value: #{@val}, prev: #{@prev&.value}, next: #{@nnext&.value}>"
    end

    def to_s
      self.inpect
    end
  end

  attr_reader :head, :tail

  def initialize(head: nil, tail: nil)
    @head = head
    @tail = tail
  end

  def <<(val)
    node = Node.new(val)
    if @head.nil?
      @head = node
      @tail = node
    else
      node.prev = @tail
      @tail.next = node
      @tail = node
    end
  end

  def inspect
    self.to_s
  end

  def to_s
    builder = StringIO.new

    cur = @head

    until cur.nil?
      builder << cur.value
      builder << ','
      cur = cur.next
    end

    builder.string
  end
end

list = List.new()

list.to_s

list << "111"

list.head.next

hi = files.length - 1

while hi > 0
  file = files[hi]
  lo = 0

  while lo < hi
    if lo.free > file.size
      # put file right after lo
      # add file size to hi - 1 free space
      # break
     end

    lo += 1
  end

  hi -= 1
end

checksum = 0

res.each.with_index { |id, i| break if id == '.'; checksum += id.to_i * i };

checksum.to_i

