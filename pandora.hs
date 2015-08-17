{-# LANGUAGE OverloadedStrings #-}
{-# LANGUAGE DeriveGeneric #-}

import Control.Monad
import Control.Applicative ((<$>))
import Data.ByteString.Char8
import Data.ByteString.UTF8 (fromString)
import Data.Aeson (FromJSON, ToJSON, decode, encode)
import Data.Set (Set)
import qualified Data.Set as Set
import qualified Data.ByteString.Lazy.Char8 as BL
import Debug.Trace
import GHC.Generics (Generic)
import System.Environment
import System.ZMQ4.Monadic
import System.Posix.Daemon
import Text.Pandoc
import Text.Printf


-- TODO use delete function
myPandocExtensions :: Set Extension
myPandocExtensions = Set.fromList
  [ Ext_footnotes
  , Ext_inline_notes
  , Ext_pandoc_title_block
  , Ext_yaml_metadata_block
  , Ext_table_captions
  , Ext_implicit_figures
  , Ext_simple_tables
  , Ext_multiline_tables
  , Ext_grid_tables
  , Ext_pipe_tables
  , Ext_citations
  , Ext_raw_tex
  , Ext_raw_html
  , Ext_tex_math_dollars
  , Ext_latex_macros
  , Ext_fenced_code_blocks
  , Ext_fenced_code_attributes
  , Ext_backtick_code_blocks
  , Ext_inline_code_attributes
  , Ext_markdown_in_html_blocks
--   , Ext_escaped_line_breaks
  , Ext_fancy_lists
  , Ext_startnum
  , Ext_definition_lists
  , Ext_example_lists
  , Ext_all_symbols_escapable
  , Ext_intraword_underscores
  , Ext_blank_before_blockquote
  , Ext_blank_before_header
  , Ext_strikeout
  , Ext_superscript
  , Ext_subscript
  , Ext_auto_identifiers
  , Ext_header_attributes
  , Ext_implicit_header_references
  , Ext_line_blocks
  ]


updateDefautlOptions :: WriterOptions -> Set Extension -> WriterOptions
updateDefautlOptions opts extensions = opts{ writerExtensions =  extensions }




data Input = Input{
                  text :: String ,
                  from :: String ,
                  to :: String
            }
             deriving (Show, Generic)

instance FromJSON Input
instance ToJSON Input


addr = "tcp://127.0.0.1:9999"

workerURL = "inproc://workers"


main :: IO ()
main = runDetached (Nothing) def $ runZMQ $ do
        -- Socket to talk to clients
        clients <- socket Router
        bind clients addr

        -- Socket to talk to workers
        workers <- socket Dealer
        bind workers workerURL

        -- using inproc (inter-thread) we expect to share the same context
        replicateM_ 200 (async worker)

        -- Connect work threads to client threads via a queue
        proxy clients workers Nothing


worker :: ZMQ z ()
worker = do
    receiver <- socket Rep
    connect receiver workerURL
    forever $ do
        unpack <$> receive receiver >>= send receiver [] . pack . utf8 . transform . unmarshalJSON . toLazyBytes


toLazyBytes :: [Char] -> BL.ByteString
toLazyBytes s = BL.pack s


unmarshalJSON :: BL.ByteString -> Maybe Input
unmarshalJSON bs = decode bs :: Maybe Input


transform :: Maybe Input -> String
-- transform a | trace ("transform: " ++ show a) False = undefined
transform a = case a of
    Nothing -> "opppps!"
    Just x -> do
        let f = from x in case f of
         "html" -> let t = to x in case t of
              "markdown" -> html2markdown $ text x
              "html" ->  text x
         "markdown" -> let t = to x in case t of
              "html" -> markdown2html $ text x
              "markdown" -> text x


utf8 :: String -> String
utf8 = unpack . fromString


html2markdown :: String -> String
html2markdown = parse (readHtml def) (writeMarkdown (updateDefautlOptions def myPandocExtensions))

markdown2html :: String -> String
markdown2html = parse (readMarkdown def) (writeHtmlString def)

parse reader writer input = unwrap $ do
 parsed <- reader input
 return $ writer parsed

-- Ignoring errors!
unwrap :: Either a String -> String
unwrap (Left _) = ""
unwrap (Right string) = string
