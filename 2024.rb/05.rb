printer_instructions = <<~TEXT
47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47
TEXT

printer_instructions = File.read('./05_input.txt')

parts = printer_instructions.split("\n\n")

order_rules = parts.first.split("\n").map { _1.split('|').map(&:to_i) }.tap { p _1 }.each_with_object({}) do |(x, y), obj| 
  obj[x] = [] unless obj.include?(x)
  obj[x] << y
end

page_updates = parts[1].split("\n").map { _1.split(',').map(&:to_i) }

page_updates.filter do |update|
  seen = Set.new

  update.none? do |page_num|
    seen << page_num
    order_rules.fetch(page_num, []).any? { seen.include?(_1) }
  end
end.sum { |valid_update| valid_update[valid_update.length / 2] }

# part two

order_rules_rev = parts.first.split("\n").map { _1.split('|').map(&:to_i) }.tap { p _1 }.each_with_object({}) do |(x, y), obj| 
  obj[y] = Set.new unless obj.include?(y)
  obj[y] << x
end

unordered = page_updates.filter do |update|
  seen = Set.new

  update.any? do |page_num|
    seen << page_num
    order_rules.fetch(page_num, []).any? { seen.include?(_1) }
  end
end

unordered.map do |xs|
  xs.sort do |a, b|
    if order_rules_rev.fetch(a, Set.new).include?(b)
      1
    else
      -1
    end
  end
end.sum { |valid_update| valid_update[valid_update.length / 2] }
