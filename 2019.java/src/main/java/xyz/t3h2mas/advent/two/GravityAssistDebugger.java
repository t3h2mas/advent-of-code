package xyz.t3h2mas.advent.two;

public class GravityAssistDebugger {
    private final IntComputer intComputer = new IntComputer();
    private final int NOUN_ADDRESS = 1;
    private final int VERB_ADDRESS = 2;

    public void recreate1202Alarm(int[] program) {
        int[] localProgram = program.clone();
        localProgram[NOUN_ADDRESS] = 12;
        localProgram[VERB_ADDRESS] = 2;

        try {
            intComputer.compute(localProgram);
        } catch (HaltSignaledException e) {
        }

        System.out.println("Recreate 1202 Alarm: Position 0 state: " + localProgram[0]);
    }

    public void findRequiredVerbAndNoun(int[] program) {
        final int NEEDLE = 19690720;

        for (int n = 0; n < 100; n++) {
            for (int v = 0; v < 100; v++) {
                int[] localProgram = program.clone();
                localProgram[NOUN_ADDRESS] = n;
                localProgram[VERB_ADDRESS] = v;

                try {
                    intComputer.compute(localProgram);
                } catch (HaltSignaledException e) {
                }

                int actual = localProgram[0];

                if (actual == NEEDLE) {
                    int calculatedCombination = 100 * n + v;
                    System.out.println("Found noun/verb combination: " + calculatedCombination);
                    return;
                }
            }
        }
        System.out.println("Failed to find noun/verb combination...");
    }
}
