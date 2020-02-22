package br.com.zup.beagle.sample.controller

import br.com.zup.beagle.core.ServerDrivenComponent
import br.com.zup.beagle.widget.core.Flex
import br.com.zup.beagle.widget.core.FlexDirection
import br.com.zup.beagle.widget.layout.Container
import br.com.zup.beagle.widget.layout.ScrollAxis
import br.com.zup.beagle.widget.layout.ScrollView
import br.com.zup.beagle.widget.ui.Text
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController

@RestController
class ScrollViewController {

    @RequestMapping("/scrollView")
    fun getPageView(): ServerDrivenComponent {
        return ScrollView(
            scrollDirection = ScrollAxis.HORIZONTAL,
            children = listOf(
                Container(
                    children = listOf(
                        Text("Texto"),
                        Text("Texto"),
                        Text("Texto"),
                        Text("Texto"),
                        Text("Texto"),
                        Text("Texto"),
                        Text("Texto"),
                        Text("Texto"),
                        Text("Texto"),
                        Text("Texto"),
                        Text("Texto"),
                        Text("Texto"),
                        Text("Texto"),
                        Text("Texto"),
                        Text("Texto"),
                        Text("Texto"),
                        Text("Texto"),
                        Text("Texto"),
                        Text("Texto"),
                        Text("Texto"),
                        Text("Texto")
                    )
                ).applyFlex(
                    flex = Flex(
                        flexDirection = FlexDirection.ROW
                    )
                )
            )
        )
    }
}