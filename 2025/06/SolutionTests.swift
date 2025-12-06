import Testing
import AOCTestSupport
@testable import Day06

@Suite("Day 06 Solutions")
struct Day06Tests {
    @Test("All test cases from tests.yaml")
    func testFromYAML() throws {
        try runAOCTests(bundle: .module, part1Solutions: part1Solutions, part2Solutions: part2Solutions)
    }
}
