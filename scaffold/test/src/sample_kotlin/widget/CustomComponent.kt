package br.com.zup.beagle.sample.widget

import br.com.zup.beagle.widget.core.ComposeComponent
import br.com.zup.beagle.widget.core.Flex
import br.com.zup.beagle.widget.core.JustifyContent
import br.com.zup.beagle.widget.Widget
import br.com.zup.beagle.widget.layout.Container
import br.com.zup.beagle.widget.ui.Button
import br.com.zup.beagle.widget.ui.Text

class CustomComponent : ComposeComponent() {
    override fun build(): Widget = Container(
        children = listOf(
            Button("Text 1"),
            Text("Text 1")
        )
    ).applyFlex(flex = Flex(
        justifyContent = JustifyContent.CENTER
    )
    )
}