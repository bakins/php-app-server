<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: opentelemetry/proto/common/v1/common.proto

namespace Opentelemetry\Proto\Common\V1;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * InstrumentationScope is a message representing the instrumentation scope information
 * such as the fully qualified name and version. 
 *
 * Generated from protobuf message <code>opentelemetry.proto.common.v1.InstrumentationScope</code>
 */
class InstrumentationScope extends \Google\Protobuf\Internal\Message
{
    /**
     * An empty instrumentation scope name means the name is unknown.
     *
     * Generated from protobuf field <code>string name = 1;</code>
     */
    protected $name = '';
    /**
     * Generated from protobuf field <code>string version = 2;</code>
     */
    protected $version = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $name
     *           An empty instrumentation scope name means the name is unknown.
     *     @type string $version
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Opentelemetry\Proto\Common\V1\Common::initOnce();
        parent::__construct($data);
    }

    /**
     * An empty instrumentation scope name means the name is unknown.
     *
     * Generated from protobuf field <code>string name = 1;</code>
     * @return string
     */
    public function getName()
    {
        return $this->name;
    }

    /**
     * An empty instrumentation scope name means the name is unknown.
     *
     * Generated from protobuf field <code>string name = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setName($var)
    {
        GPBUtil::checkString($var, True);
        $this->name = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string version = 2;</code>
     * @return string
     */
    public function getVersion()
    {
        return $this->version;
    }

    /**
     * Generated from protobuf field <code>string version = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setVersion($var)
    {
        GPBUtil::checkString($var, True);
        $this->version = $var;

        return $this;
    }

}

