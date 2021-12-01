module AdventIO
  def self.as_lines(file)
    file.read.split('\n').filter { |line| line.length > 0 }
  end
end
