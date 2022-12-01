package xyz.t3h2mas.advent.five;

import xyz.t3h2mas.advent.two.GravityAssistDebugger;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;

public class BootLoader {
    private static final String inputFilePath = "src/main/java/xyz/t3h2mas/advent/five/input.txt";
    public static void main(String[] args) throws IOException {
        String contents = new String(Files.readAllBytes(Paths.get(inputFilePath)));

        int[] program = makeMachineReadable(contents);

        IntComputer computer = new IntComputer();

        System.out.println("[*] running thing");
        try {
            computer.compute(program);
        } catch (HaltSignaledException e) {
            System.out.println("[*] Done...");
        }
    }

    private static int[] makeMachineReadable(String program) {
        String[] chunks = program.split(",");
        int[] result = new int[chunks.length];
        for (int i=0; i < chunks.length; i++) {
            result[i] = Integer.parseInt(chunks[i]);
        }

        return result;
    }
}
