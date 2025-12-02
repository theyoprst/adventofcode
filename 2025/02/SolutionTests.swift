import Testing
import AOCTestSupport
@testable import Day02

@Suite("Day 02 Solutions")
struct Day02Tests {
    @Test("All test cases from tests.yaml")
    func testFromYAML() throws {
        try runAOCTests(bundle: .module, part1Solutions: part1Solutions, part2Solutions: part2Solutions)
    }
}
