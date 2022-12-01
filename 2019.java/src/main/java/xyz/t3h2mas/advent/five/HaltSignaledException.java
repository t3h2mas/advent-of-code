package xyz.t3h2mas.advent.five;

public class HaltSignaledException extends RuntimeException {
    public HaltSignaledException(String message) {
        super(message);
    }
}
