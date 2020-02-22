package br.com.zup.beagle.sample.controller

import br.com.zup.beagle.action.Navigate
import br.com.zup.beagle.action.NavigationType
import br.com.zup.beagle.annotation.BeaglePreview
import br.com.zup.beagle.widget.core.Flex
import br.com.zup.beagle.widget.core.JustifyContent
import br.com.zup.beagle.widget.layout.Container
import br.com.zup.beagle.widget.layout.NavigationBar
import br.com.zup.beagle.widget.layout.NavigationBarItem
import br.com.zup.beagle.widget.layout.Screen
import br.com.zup.beagle.widget.ui.Button
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController

@RestController
class SampleViewController {

    @BeaglePreview
    @RequestMapping("/sample")
    fun getSampleView(): Screen {

        return Screen(
            navigationBar = NavigationBar(
                title = "Sample Bar",
                style = "Style.Default",
                navigationBarItems = listOf(
                    NavigationBarItem(
                        text = "First",
                        image = "delete",
                        action = Navigate(
                            type = NavigationType.POP_VIEW,
                            path = "http://localhost:8080/sample"
                        )
                    ),
                    NavigationBarItem(
                        text = "Second",
                        image = "question",
                        action = Navigate(
                            type = NavigationType.POP_VIEW
                        )
                    )
                )
            ),
            content = Container(
                children = listOf(
                    Button("")
                )
            ).applyFlex(Flex(justifyContent = JustifyContent.CENTER))
        )
    }
}
