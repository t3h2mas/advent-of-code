package xyz.t3h2mas.advent.one;

import xyz.t3h2mas.advent.one.domain.*;
import xyz.t3h2mas.advent.one.domain.Module;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

public final class FuelCounterUpperLauncher {
    private static final String inputFilePath = "src/main/java/xyz/t3h2mas/advent/one/input.txt";

    public static void main(String[] args) throws FileNotFoundException {
        File input = new File(FuelCounterUpperLauncher.inputFilePath);
        Scanner scanner = new Scanner(input);

        SpaceCraft spaceCraft = new SpaceCraft();
        while (scanner.hasNextInt()) {
            spaceCraft.addModule(new Module(scanner.nextInt()));
        }


        runSolutionOne(spaceCraft);
        runSolutionTwo(spaceCraft);
    }

    public static void runSolutionOne(SpaceCraft spaceCraft) {
        System.out.println("Solution One:");

        FuelStrategy fuelStrategy = new FuelForModules();
        FuelCounterUpper fuelCounterUpper = new FuelCounterUpper(fuelStrategy);

        System.out.println(fuelCounterUpper.calculateFuelRequirements(spaceCraft));
    }

    private static void runSolutionTwo(SpaceCraft spaceCraft) {
        System.out.println("Solution Two:");

        FuelStrategy fuelStrategy = new FuelForModulesIncludingFuel();
        FuelCounterUpper fuelCounterUpper = new FuelCounterUpper(fuelStrategy);

        System.out.println(fuelCounterUpper.calculateFuelRequirements(spaceCraft));
    }
}
