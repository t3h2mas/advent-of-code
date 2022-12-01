require_relative './lib/chunk_validator'

corrupt_scores = {
  ')' => 3,
  ']' => 57,
  '}' => 1197,
  '>' => 25137
}

autocomplete_scores = {
  ')' => 1,
  ']' => 2,
  '}' => 3,
  '>' => 4
}

chunks = File.read('./input.txt').split("\n").filter { |l| !l.empty? }
#chunks = File.read('./input2.txt').split("\n").filter { |l| !l.empty? }

corrupted_score = 0
chunks.each do |chunk|
  valid = ChunkValidator.validate(chunk)

  next unless valid[:result] == :corrupted

  corrupted_score += corrupt_scores[valid[:char]]
end

puts "Corrupted score: #{corrupted_score}"

completed_scores = []

chunks.each do |chunk|
  valid = ChunkValidator.validate(chunk)

  next unless valid[:result] == :incomplete

  autocomplete_score = 0

  valid[:completers].each do |c|
    autocomplete_score *= 5
    autocomplete_score += autocomplete_scores[c]
  end

  completed_scores << autocomplete_score
end

completed_scores.sort!

puts "Autocomplete score: #{completed_scores[completed_scores.size / 2]}"
