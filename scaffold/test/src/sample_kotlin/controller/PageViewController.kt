package br.com.zup.beagle.sample.controller

import br.com.zup.beagle.widget.core.Alignment
import br.com.zup.beagle.widget.core.Flex
import br.com.zup.beagle.widget.core.JustifyContent
import br.com.zup.beagle.widget.Widget
import br.com.zup.beagle.widget.layout.PageView
import br.com.zup.beagle.widget.pager.PageIndicator
import br.com.zup.beagle.widget.ui.Text
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController

@RestController
class PageViewController {

    @RequestMapping("/pageView")
    fun getPageView(): Widget {
        return PageView(
            pageIndicator = PageIndicator(
                selectedColor = "#FFFFFF",
                unselectedColor = "#888888"
            ),
            pages = listOf(
                Text("Page 1").applyFlex(
                    flex = Flex(
                        justifyContent = JustifyContent.CENTER,
                        alignItems = Alignment.CENTER
                    )
                ),
                Text("Page 2").applyFlex(
                    flex = Flex(
                        justifyContent = JustifyContent.CENTER,
                        alignItems = Alignment.CENTER
                    )
                ),
                Text("Page 3").applyFlex(
                    flex = Flex(
                        justifyContent = JustifyContent.CENTER,
                        alignItems = Alignment.CENTER
                    )
                ),
                Text("Page 4").applyFlex(
                    flex = Flex(
                        justifyContent = JustifyContent.CENTER,
                        alignItems = Alignment.CENTER
                    )
                ),
                Text("Page 5").applyFlex(
                    flex = Flex(
                        justifyContent = JustifyContent.CENTER,
                        alignItems = Alignment.CENTER
                    )
                )
            )
        )
    }
}