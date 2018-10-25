package utils

import (
    "github.com/google/logger"
)

func loggerTest() {
    logger.Info("loggerTest")
    logger.Warning("loggerTest")
    logger.Error("loggerTest")
}

//# -*- coding:utf-8 -*-
//from __future__ import absolute_import
//
//import logging
//import os
//import re
//import tempfile
//import time
//from urllib.parse import urlparse
//from urllib.request import urlretrieve
//
//import qreader
//import requests
//from bs4 import BeautifulSoup
//
//from basic_tools import create_guid, get_md5
//from basic_tools import smart_decoder
//from shadowsocks_tools.account_object import SSAccountObj
//
//logger = logging.getLogger(__name__)
//
//
//def simple_request_get(url):
//    return requests.get(
//        url,
//        headers={
//            'user-agent': 'Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.2 (KHTML, like Gecko)' +
//                          ' Chrome/22.0.1216.0 Safari/537.2'
//        },
//        timeout=15).content
//
//
//def proxy_request_get(url, proxy_addr, proxy_port):
//    socks_url = "socks5://{}:{}".format(proxy_addr, proxy_port)
//    proxies = {'http': socks_url, 'https': socks_url}
//    return requests.get(
//        url,
//        headers={
//            'user-agent': 'Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.2 (KHTML, like Gecko)' +
//                          ' Chrome/22.0.1216.0 Safari/537.2'
//        },
//        proxies=proxies,
//        timeout=15).content
//
//
//def request_get_by_web_api(url, ):
//    default_headers = {
//        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36",
//        "DNT": "1",
//    }
//
//    def get_response_from_codebeautify(_url):
//        """
//        https://codebeautify.org/source-code-viewer
//        :param _url:
//        :return:
//        """
//        headers = dict(default_headers)
//        headers.update({
//            "Accept": "text/plain, */*; q=0.01",
//            "Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
//            "Origin": "https://codebeautify.org",
//            "Referer": "https://codebeautify.org/source-code-viewer",
//        })
//        return requests.post(
//            "https://codebeautify.com/URLService",
//            data={"path": _url},
//            headers=headers,
//            timeout=15).content
//
//    return get_response_from_codebeautify(url)
//
//
//def strong_request_get(url, proxy_addr="127.0.0.1", proxy_port=1080):
//    try:
//        return simple_request_get(url)
//    except Exception as e:
//        logger.error(e)
//    try:
//        return proxy_request_get(url, proxy_addr, proxy_port)
//    except Exception as e:
//        logger.error(e)
//        return request_get_by_web_api(url)
//
//
//def get_domain(url):
//    """
//
//    :rtype: str
//    :type url: str
//    """
//    parsed_uri = urlparse(url)
//    return '{uri.scheme}://{uri.netloc}/'.format(uri=parsed_uri)
//
//
//def read_qrcode_from_image_file(image_file_path, remove_file=True):
//    """
//    :type remove_file: bool
//    :type image_file_path: str
//    :rtype: str
//    """
//    # read qr image
//    data = qreader.read(image_file_path)
//    if remove_file and os.path.exists(image_file_path):
//        os.remove(image_file_path)
//
//    return data
//
//
//def read_qrcode_from_base64(html_base64_image):
//    """
//    html_base64_image, example:
//        data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAUgAAAFIAQAAAAAab1qeAA
//        AACXBIWXMAAAsSAAALEgHS3X78AAABa0lEQVRoge3VwW7DMAwDUP3/T3NoRUp2NmBSzk
//        yH1LGed1AENjC9wtLS0tLyKSOv72PgWB0Vy4Vka+O7+hRYxXG3XMhPhbvV+NyriuVa8s
//        7+x/0/LNcSyglwoC1fSnn+IepzI8uR5O/TH5/nL5flQNbFTVTunjXLqWTfKx5yoK/gtd
//        zJJDRRoVvHLReSzWeTlQSBTl/LlYwOA021nmC5ltHX+Sqy+jtDLP+RbLdarFHG49FyLl
//        Vh8znBRLBcSnBu84uiz1vuZVWPNxANLReShKXqeJ+1XMmusdXsPo5osFxIZkJ0vxvnhu
//        VCZjDUxjnPEpZzGVEc1/AqKix3Uvuo/gfHuk9aLmT2/F6ofIy45VDeEdCn1H7LjQSuY6
//        Fz2fczGSwnMtTjYLiKgBuWO6llqOkZEMqFsFzKCEUBat1vBJbvJNdZYTYQWr6T0PhCg2
//        25l/R1h7K3qeVcXvvKCG1qsi2ncnRZWlpaWp7XDyBajl7ewO8kAAAAAElFTkSuQmCC
//    :rtype: str
//    :type html_base64_image: str
//    """
//    tmp_file_name = os.path.join(
//        tempfile.gettempdir(),
//        "{}.{}".format(get_md5(create_guid()), html_base64_image[11:html_base64_image.find(";")])
//    )
//
//    # save qr image
//    urlretrieve(html_base64_image, filename=tmp_file_name)
//
//    return read_qrcode_from_image_file(tmp_file_name, remove_file=True)
//
//
//re_base64_image = re.compile(r'data:image/[a-zA-Z]+;base64,[a-zA-Z0-9/+]+={0,2}')
//
//
//def list_ss_url_from_base64(content):
//    """
//    通过正则方法解析网页内容中的base64 image内容
//    :rtype: list of str
//    :type content: str
//    """
//    base64_list = re_base64_image.findall(content)
//    ss_url_list = []
//
//    for base64_content in base64_list:
//        try:
//            ss_url = read_qrcode_from_base64(base64_content)
//            if ss_url:
//                ss_url_list.append(ss_url)
//        except Exception as e:
//            logger.error(e)
//
//    return [ss_url for ss_url in ss_url_list if ss_url.startswith("ss")]
//
//
//re_ss_url = re.compile(r'ss://[a-zA-Z0-9]+')
//re_vmess_url = re.compile(r'vmess://[a-zA-Z0-9]+')
//
//
//def list_ss_url_from_text(content):
//    """
//    通过正则方法解析网页内容中的base64 image内容
//    :rtype: list of str
//    :type content: str
//    """
//    ss_url_list = re_ss_url.findall(content)
//    vmess_url_list = re_vmess_url.findall(content)
//
//    def _is_vmess(_ss_url):
//        if not vmess_url_list:
//            return False
//
//        for _vmess in vmess_url_list:
//            if _vmess[3:] == _ss_url:
//                return True
//
//        return False
//
//    return [ss_url for ss_url in ss_url_list if not _is_vmess(ss_url)]
//
//
//def list_ss_url_from_image(content, domain):
//    """
//    解析网页内容中的图片二维码
//    domain example: https://go.freess.org
//    :type domain: str
//    :rtype: list of str
//    :type content: str
//    """
//    # find all images
//    soup = BeautifulSoup(content)
//    images = []
//
//    # for img in soup.findAll('img'):
//    #     src = img.get('src')
//    #     if src:
//    #         images.append(src)
//
//    # 根据qr必须显示在屏幕中间，推测qr不在img标签内; 又根据qr为黑白色，推测图片格式为png
//    for a in soup.find_all('a', href=True):
//        try:
//            image = a['href']
//            if image and image.endswith("png"):
//                images.append(image)
//        except Exception as e:
//            logger.error(e)
//
//    # list all image urls
//    image_urls = []
//    for image in images:
//        if image.startswith("http"):
//            image_urls.append(image)
//        elif image.startswith("/"):
//            image_urls.append(domain + image)
//        else:
//            image_urls.append(domain + "/" + image)
//
//    # read data from images
//    ss_url_list = []
//    tmp_dir = tempfile.gettempdir()
//    for image in image_urls:
//        try:
//            print(image)
//            tmp_file_name = os.path.join(
//                tmp_dir,
//                "{}.{}".format(get_md5(create_guid()), image.split(".")[-1])
//            )
//
//            # save qr image
//            file_content = strong_request_get(image)
//            with open(tmp_file_name, "wb") as f:
//                f.write(file_content)
//
//            ss_url = read_qrcode_from_image_file(tmp_file_name, remove_file=True)
//
//            if ss_url:
//                ss_url_list.append(ss_url)
//        except Exception as e:
//            logger.error(e)
//
//    return [ss_url for ss_url in ss_url_list if ss_url.startswith("ss")]
//
//
//def test_list_ss_url_from_base64():
//    url_list = ["https://free-ss.site/"]
//
//    ss_url_list = []
//    for url in url_list:
//        try:
//            text = smart_decoder(strong_request_get(url))
//            _ss_url_list = list_ss_url_from_base64(text)
//            if _ss_url_list:
//                ss_url_list.extend(_ss_url_list)
//            else:
//                logger.warning(text)
//
//        except Exception as e:
//            logger.error(e)
//
//    [logger.info(ss_url) for ss_url in ss_url_list]
//
//
//def test_list_ss_url_from_image():
//    url_list = ["https://us.ishadowx.net"]
//
//    ss_url_list = []
//    for url in url_list:
//        try:
//            text = smart_decoder(strong_request_get(url))
//            _ss_url_list = list_ss_url_from_image(text, url)
//            if _ss_url_list:
//                ss_url_list.extend(_ss_url_list)
//            else:
//                logger.warning(text)
//
//        except Exception as e:
//            logger.error(e)
//
//    [logger.info(ss_url) for ss_url in ss_url_list]
//
//
//def parse_html(content, domain=None):
//    """
//
//    :type domain: str|None
//    :rtype: list of SSAccountObj
//    :type content: str
//    """
//    account_list = []
//
//    # parse ss account from text
//    # todo not finish
//
//    # parse ss_url from text
//    all_ss_url = list_ss_url_from_text(content)
//
//    # parse_ss_url from base64
//    all_ss_url.extend(list_ss_url_from_base64(content))
//
//    # parse_ss_url from image
//    if not account_list and domain:
//        all_ss_url.extend(list_ss_url_from_image(content, domain))
//
//    # ss_account_obj from ss_url
//    for ss_url in all_ss_url:
//        try:
//            account_list.append(SSAccountObj.from_ss_protocol(ss_url))
//        except Exception as e:
//            logger.error(e)
//
//    # 去重
//    account_dict = {account.hash_code: account for account in account_list}
//    return account_dict.values()
//
//
//def list_ss_account_of_url(url):
//    """
//
//    :type url: str
//    :rtype: list of SSAccountObj
//    """
//    try:
//        text = smart_decoder(strong_request_get(url))
//        return parse_html(text, domain=get_domain(url))
//    except Exception as e:
//        logger.error(e)
//        return []
//
//
//re_match_url = re.compile(r'http[s]?://(?:[a-zA-Z]|[0-9]|[$-_@.&+]|[!*\(\),]|(?:%[0-9a-fA-F][0-9a-fA-F]))+')
//
//
//def list_target_url_from_setting():
//    """
//
//    :rtype: list of url
//    """
//
//    def list_target_url_from_default_setting():
//        all_url_list = [
//            "https://us.ishadowx.net/",
//            "https://ss.freess.org/",
//            "https://free-ss.site/",
//            "https://doub.io/sszhfx",
//        ]
//
//        # url_share
//        url = "https://raw.githubusercontent.com/loremwalker/WebSiteUseful/master" + \
//              "/%E7%A7%91%E5%AD%A6%E4%B8%8A%E7%BD%91/SS%26%26SSR%26v2ray%E5%88%86%E4%BA%AB.md"
//
//        try:
//            content = smart_decoder(strong_request_get(url))
//            all_url_list.extend(re_match_url.findall(content))
//        except Exception as e:
//            logger.error(e)
//
//        return list(set(all_url_list))
//
//    from shadowsocks_tools.ss_local_tool.ss_client_controller import get_config_file_manager
//    config_file_manager = get_config_file_manager()
//
//    urls_from_config = config_file_manager.target_urls if config_file_manager.target_urls else []
//
//    if config_file_manager.only_use_target_urls_in_config_file is False:
//        urls_from_config.extend(list_target_url_from_default_setting())
//
//    return list(set(urls_from_config))
//
//
//if __name__ == '__main__':
//    from shadowsocks_tools.ss_local_tool.ss_client_controller import init_config_file_manager, get_config_file_manager
//
//    target_url_list = ["https://us.ishadowx.net/", "https://ss.freess.org/"]
//
//    config_file_path = "E://Program/Shadowsocks-4.0.9/gui-config.json"
//
//    init_config_file_manager(config_file=config_file_path, use_ha_mode=True)
//    editor = get_config_file_manager()
//
//    # first
//    _account = editor.list_account_of_current_config_file()[0]
//    logger.info("{} {}".format("Valid" if _account.is_valid else "Invalid", _account.to_string()))
//
//    for target_url in target_url_list:
//        logger.info("from {} ==>".format(target_url))
//        account_list = list_ss_account_of_url(target_url)
//        for _account in account_list:
//            logger.info("{} {}".format("Valid" if _account.is_valid else "Invalid", _account.to_string()))
//
//        logger.info("\n")
//
//    # twice
//    _account = editor.list_account_of_current_config_file()[0]
//    logger.info("{} {}".format("Valid" if _account.is_valid else "Invalid", _account.to_string()))
//
//    for target_url in target_url_list:
//        logger.info("from {} ==>".format(target_url))
//        account_list = list_ss_account_of_url(target_url)
//        for _account in account_list:
//            logger.info("{} {}".format("Valid" if _account.is_valid else "Invalid", _account.to_string()))
//
//        logger.info("\n")
//
//    # loop
//    while True:
//        time.sleep(1)
