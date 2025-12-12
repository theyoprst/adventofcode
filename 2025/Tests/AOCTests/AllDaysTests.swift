import Testing
@testable import aoc

// TDOO: can we "generate" tests for all existing days?
@Suite("Advent of Code 2025")
struct AllDaysTests {
    @Test("Day 01")
    func testDay01() throws {
        try runAOCTests(bundle: .module, day: Day01())
    }

    @Test("Day 02")
    func testDay02() throws {
        try runAOCTests(bundle: .module, day: Day02())
    }

    @Test("Day 03")
    func testDay03() throws {
        try runAOCTests(bundle: .module, day: Day03())
    }

    @Test("Day 04")
    func testDay04() throws {
        try runAOCTests(bundle: .module, day: Day04())
    }

    @Test("Day 05")
    func testDay05() throws {
        try runAOCTests(bundle: .module, day: Day05())
    }

    @Test("Day 06")
    func testDay06() throws {
        try runAOCTests(bundle: .module, day: Day06())
    }

    @Test("Day 07")
    func testDay07() throws {
        try runAOCTests(bundle: .module, day: Day07())
    }

    @Test("Day 08")
    func testDay08() throws {
        try runAOCTests(bundle: .module, day: Day08())
    }

    @Test("Day 09")
    func testDay09() throws {
        try runAOCTests(bundle: .module, day: Day09())
    }

    @Test("Day 10")
    func testDay10() throws {
        try runAOCTests(bundle: .module, day: Day10())
    }

    @Test("Day 11")
    func testDay11() throws {
        try runAOCTests(bundle: .module, day: Day11())
    }

    @Test("Day 12")
    func testDay12() throws {
        try runAOCTests(bundle: .module, day: Day12())
    }
}
