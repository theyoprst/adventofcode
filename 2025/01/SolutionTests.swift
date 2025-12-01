import Testing
import AOCTestSupport
@testable import Day01

@Suite("Day 01 Solutions")
struct Day01Tests {
    @Test("All test cases from tests.yaml")
    func testFromYAML() throws {
        try runAOCTests(bundle: .module, part1Solutions: part1Solutions, part2Solutions: part2Solutions)
    }
}
