require "./clients/ruby/pandora"

sampleHTML = """
        <ul>
        <li>
        你好
        </li>
        <li>
        你妹
        </li>
        </ul>
"""

include Pandora
md = to_markdown sampleHTML
puts md


sampleMD = """
- 你好
- 你妹
"""

html = to_html sampleMD
puts html
