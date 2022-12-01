require_relative './lib/submarine'

input = File.read('./input.txt')

commands = input.split("\n").filter { |line| !line.nil? && line != "" }.map do |line|
  parts = line.split(" ")
  direction = parts[0].to_sym
  count = parts[1].to_i

  { direction: direction, count: count }
end

directional_submarine = Submarine.new(DirectionalNavigation.new)
commands.each { |command| directional_submarine.move(command[:direction], command[:count]) }
puts "Part one: #{directional_submarine.depth * directional_submarine.horizontal_position}"


aiming_submarine = Submarine.new(AimNavigation.new)
commands.each { |command| aiming_submarine.move(command[:direction], command[:count]) }
puts "Part two: #{aiming_submarine.depth * aiming_submarine.horizontal_position}"
