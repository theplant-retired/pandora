require 'rubygems'
require 'ffi-rzmq'
require 'json'

module Pandora
  def to_markdown text
    Converter.new.to_markdown text
  end

  def to_html text
    Converter.new.to_html text
  end

  class Converter
    FLAG_HTML = "html"
    FLAG_MARKDOWN = "markdown"
    def initialize
      context = ZMQ::Context.new(1)
      @requester = context.socket(ZMQ::REQ)
      @requester.connect("tcp://127.0.0.1:9999")
    end

    def to_markdown text 
      convert text, FLAG_HTML, FLAG_MARKDOWN
    end

    def to_html text
      convert text, FLAG_MARKDOWN, FLAG_HTML
    end

    private
    def convert(text, from, to)
      obj = {text: text, from: from, to: to}
      msg = obj.to_json
      @requester.send_string msg
      reply = ''
      @requester.recv_string(reply)
      reply
    end

  end
end
