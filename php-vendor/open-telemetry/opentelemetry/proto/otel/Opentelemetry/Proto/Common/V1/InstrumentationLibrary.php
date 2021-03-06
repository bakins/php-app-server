<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: opentelemetry/proto/common/v1/common.proto

namespace Opentelemetry\Proto\Common\V1;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * InstrumentationLibrary is a message representing the instrumentation library information
 * such as the fully qualified name and version.
 * InstrumentationLibrary is wire-compatible with InstrumentationScope for binary
 * Protobuf format.
 * This message is deprecated and will be removed on June 15, 2022.
 *
 * Generated from protobuf message <code>opentelemetry.proto.common.v1.InstrumentationLibrary</code>
 */
class InstrumentationLibrary extends \Google\Protobuf\Internal\Message
{
    /**
     * An empty instrumentation library name means the name is unknown.
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
     *           An empty instrumentation library name means the name is unknown.
     *     @type string $version
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Opentelemetry\Proto\Common\V1\Common::initOnce();
        parent::__construct($data);
    }

    /**
     * An empty instrumentation library name means the name is unknown.
     *
     * Generated from protobuf field <code>string name = 1;</code>
     * @return string
     */
    public function getName()
    {
        return $this->name;
    }

    /**
     * An empty instrumentation library name means the name is unknown.
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

