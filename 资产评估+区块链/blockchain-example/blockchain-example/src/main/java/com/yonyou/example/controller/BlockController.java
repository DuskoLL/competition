package com.yonyou.example.controller;

import com.yonyou.example.service.BlockchainOperateService;
import org.springframework.web.bind.annotation.*;

import javax.annotation.Resource;

/**
 * @ClassName blockController
 * @Description TODO
 * @Author 孙世江
 * @Data 2020/7/28 0028 下午 9:37
 * @Version 1.0
 **/
@RestController
@RequestMapping(value = "/block")
public class BlockController {

    @Resource
    BlockchainOperateService blockchainOperateService;

    @RequestMapping("invoke/{function}")
    public Object set(@PathVariable String function, @RequestParam("args") String... args) {
        Object response = null;
        try {
            response = blockchainOperateService.set(function, args);
        } catch (Exception e) {
            e.printStackTrace();
        }
        return response;
    }

    @GetMapping("query/{function}/{key}")
    public String get(@PathVariable String function, @PathVariable String key) {
        String response = null;
        try {
            response = blockchainOperateService.get(function, key);
        } catch (Exception e) {
            e.printStackTrace();
        }
        return response;
    }
}
