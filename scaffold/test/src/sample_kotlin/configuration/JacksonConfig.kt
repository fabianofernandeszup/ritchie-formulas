package br.com.zup.beagle.sample.configuration

import br.com.zup.beagle.serialization.jackson.BeagleActionSerializer
import br.com.zup.beagle.serialization.jackson.BeagleScreenSerializer
import br.com.zup.beagle.serialization.jackson.BeagleComponentSerializer
import com.fasterxml.jackson.annotation.JsonInclude
import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration
import org.springframework.http.converter.json.Jackson2ObjectMapperBuilder

@Configuration
open class JacksonConfig {

    @Bean
    open fun jacksonBuilder(): Jackson2ObjectMapperBuilder {
        return Jackson2ObjectMapperBuilder()
            .serializationInclusion(JsonInclude.Include.NON_NULL)
            .serializers(BeagleComponentSerializer())
            .serializers(BeagleActionSerializer())
            .serializers(BeagleScreenSerializer())
    }
}