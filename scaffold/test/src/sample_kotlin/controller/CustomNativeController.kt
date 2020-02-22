package br.com.zup.beagle.sample.controller

import br.com.zup.beagle.sample.widget.CustomNativeWidget
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController

@RestController
@RequestMapping("/custom/native")
class CustomNativeController {
    @GetMapping
    fun getCustomNativeWidget() = CustomNativeWidget()
}
