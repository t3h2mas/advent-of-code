package xyz.t3h2mas.advent.two;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;

public class BootLoader {
    private static final String inputFilePath = "src/main/java/xyz/t3h2mas/advent/two/input.txt";
    public static void main(String[] args) throws IOException {
        String contents = new String(Files.readAllBytes(Paths.get(inputFilePath)));
        GravityAssistDebugger debugger = new GravityAssistDebugger();

        int[] program = makeMachineReadable(contents);

        System.out.println("[*] Recreating 1201 Alarm");
        debugger.recreate1202Alarm(program);

        System.out.println("[*] Searching for verb/noun input combination");
        debugger.findRequiredVerbAndNoun(program);
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
