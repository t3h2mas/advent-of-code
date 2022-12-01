package xyz.t3h2mas.advent.four;

public class Range {
    private int start;
    private int stop;

    public Range(int start, int stop) {
        this.start = start;
        this.stop = stop;
    }

    public int getStart() {
        return start;
    }

    public int getStop() {
        return stop;
    }

    public static Range from(String input) {
        String[] chunks = input.split("-");

        return new Range(
                Integer.parseInt(chunks[0]),
                Integer.parseInt(chunks[1])
        );
    }
}
