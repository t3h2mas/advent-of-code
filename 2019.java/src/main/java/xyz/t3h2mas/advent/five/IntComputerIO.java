package xyz.t3h2mas.advent.five;

import java.io.InputStream;
import java.io.PrintStream;
import java.util.Scanner;

public class IntComputerIO {
    private final Scanner scanner;
    private final PrintStream out;

    public IntComputerIO(InputStream in, PrintStream out) {
        this.scanner = new Scanner(in);
        this.out = out;
    }

    public int readInt() {
        System.out.print("input: ");
        return scanner.nextInt();
    }

    public void printInt(int n) {
        this.out.println(n);
    }
}
