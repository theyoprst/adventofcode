import Testing
import AOCTestSupport
@testable import Day03

@Suite("Day 03 Solutions")
struct Day03Tests {
    @Test("All test cases from tests.yaml")
    func testFromYAML() throws {
        try runAOCTests(bundle: .module, part1Solutions: part1Solutions, part2Solutions: part2Solutions)
    }
}
