package xyz.t3h2mas.advent.three;

public class ManhattanDistance {
    // Brought to you by https://stackoverflow.com/a/8224516/812076
    public static int calculate(int x0, int x1, int y0, int y1) {
       return Math.abs(x1-x0) + Math.abs(y1-y0);
    }
}
