require 'set'

class ChunkValidator
  def initialize
    @close_to_open = {
      '>' => '<',
      ']' => '[',
      ')' => '(',
      '}' => '{'
    }
    @open_to_close = @close_to_open.invert

    @opened = []
  end

  def validate(chunk)
    chunk.each_char do |c|
      if open?(c)
        @opened << c
        next
      end

      last_opener = @opened.pop

      return { result: :corrupted, char: c } if last_opener != @close_to_open[c]
    end

    return { result: :incomplete, chars: @opened, completers: completers } unless @opened.empty?

    { result: :valid }
  end

  def self.validate(chunk)
    new.validate(chunk)
  end

  private

  def open?(c)
    @open_to_close.include?(c)
  end

  def completers
    @opened.reverse.map do |c|
      @open_to_close[c]
    end
  end
end
