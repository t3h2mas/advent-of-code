package xyz.t3h2mas.advent.one.domain;

import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class FuelForModulesShould {
    private final int INPUT_MASS = 14;
    private final int EXPECTED_REQUIRED_FUEL = 2;

    @Test
    public void calculate_fuel_required_for_modules_of_a_spacecraft() {
        FuelForModules sut = new FuelForModules();

        SpaceCraft spaceCraft = new SpaceCraft();
        spaceCraft.addModule(new Module(INPUT_MASS));

        assertEquals(sut.fuelRequiredBy(spaceCraft), EXPECTED_REQUIRED_FUEL);
    }
}