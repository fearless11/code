# -*- coding:utf8 -*-
from selenium import webdriver
from selenium.webdriver.common.by import By
from time import time, sleep
import os
import datetime
import threading
import requests
from selenium.webdriver.common.desired_capabilities import DesiredCapabilities

# 需求: 抓取grafana监控页面，发送给企业微信，实现监控自动巡检

pic_path = "./pic/"
pic_name = datetime.datetime.now().strftime('%Y%m%d-%H%M') + ".png"

# 与企业微信互通的API
pic_url = "http://127.0.0.1:5566/api/v1/pic"
interurl = "127.0.0.1:3000"
interval = "3h"

# 监控面板
urls = {
    "proxy1" : "http://%s/m/dashboard/db/proxy?orgId=1&panelId=1&fullscreen&from=now-%s&to=now" % (interurl,interval),
    "proxy3" : "http://%s/m/dashboard/db/proxy?orgId=1&panelId=3&fullscreen&from=now-%s&to=now" % (interurl,interval),
    "proxy4" : "http://%s/m/dashboard/db/proxy?orgId=1&panelId=4&fullscreen&from=now-%s&to=now" % (interurl,interval),
    "proxy8" : "http://%s/m/dashboard/db/proxy?orgId=1&panelId=8&fullscreen&from=now-%s&to=now" % (interurl,interval),
    "proxy10" : "http://%s/m/dashboard/db/proxy?orgId=1&panelId=10&fullscreen&from=now-%s&to=now" % (interurl,interval), 
    "biz7" : "http://%s/m/dashboard/db/biz?orgId=1&panelId=7&fullscreen&from=now-%s&to=now" % (interurl,interval), 
}

def send_pic(pic_file):
    if not os.path.exists(pic_file):
        raise Exception("alarm Img file not exist")

    f = open(pic_file, 'rb')
    files = {"file": f}
    response = requests.post(pic_url, files=files)

    try:
        res = response.json()
        if "success" in res:
            return True
        else:
            return False
    except Exception as e:
        print str(e)
        return False


def get_urls():
    allUrls = []
    for k in urls:
        url = urls[k] + "@%s%s-%s" % (pic_path,k,pic_name)
        allUrls.append(url)
    return allUrls

	
class spier(threading.Thread):
    def __init__(self,url,pic_name):
        threading.Thread.__init__(self)
        self.url = url
        self.pic_name = pic_name

    def run(self):
        # 无头浏览器
        # wd = webdriver.PhantomJS("e:\\python\\mspier\\bin\\phantomjs.exe")
        wd = webdriver.PhantomJS("/usr/local/bin/phantomjs")

        wd.set_window_size(1900, 700)    
        wd.get(self.url)
        s = time()
        while True:
            if  time() -s > 10:
                break
            sleep(1)
        e = time()
        print("wait: %s" % (e-s),self.pic_name)
        wd.get_screenshot_as_file(self.pic_name)
        send_pic(self.pic_name)
        wd.quit()

		
if __name__ == "__main__":   
    all_urls = get_urls()
    # 线程抓取监控图
    for singe_url in all_urls:
        url = singe_url.split('@')[0]
        pic_name =  singe_url.split('@')[1]
        t = spier(url,pic_name)
        t.start()