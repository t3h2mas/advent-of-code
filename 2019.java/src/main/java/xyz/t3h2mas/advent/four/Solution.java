package xyz.t3h2mas.advent.four;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class Solution {
    private static int[] toDigits(int n) {
        return Arrays.stream(String.valueOf(n).split(""))
                .mapToInt(Integer::parseInt)
                .toArray();
    }
    private static boolean validateOne(int n) {
        int[] digits = toDigits(n);
        boolean hasDoubles = false;

        for (int i = 0; i < digits.length - 1; i++) {
            int j = i + 1;
            if (digits[i] > digits[j]) {
                return false;
            }

            if (digits[i] == digits[j]) {
                hasDoubles = true;
            }
        }

        return hasDoubles ;
    }

    public static List<Integer> partOne(Range range) {
        List<Integer> matches = new ArrayList<>();
        for (int i = range.getStart(); i < range.getStop(); i++) {
            if (Solution.validateOne(i)) {
                matches.add(i);
            }
        }

        return matches;
    }

    public static List<Integer> partTwo(Range range) {


        // 111122 meets the criteria (even though 1 is
        // repeated more than twice, it still contains a double 22).
        List<Integer> contenders = Solution.partOne(range);
        List<Integer> matches = new ArrayList<>();

        Pattern pattern = Pattern.compile("(\\d)\\1+");

        for (Integer contender : contenders) {
            boolean hasDouble = false;
            Matcher matcher = pattern.matcher(String.valueOf(contender));
            while (matcher.find()) {
               if (matcher.group().length() == 2) {
                   hasDouble = true;
                   break;
               }
            }

            if (hasDouble) {
                matches.add(contender);
            }
        }

        return matches;
    }

    public static void main(String[] cheese) {
        final String input = "231832-767346";
        final Range range = Range.from(input);
        System.out.println("Part one: " + Solution.partOne(range).size());
        System.out.println("Part two: " + Solution.partTwo(range).size());
    }
}
