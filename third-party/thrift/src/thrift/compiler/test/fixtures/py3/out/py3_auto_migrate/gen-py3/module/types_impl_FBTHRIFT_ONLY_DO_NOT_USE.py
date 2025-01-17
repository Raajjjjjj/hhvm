#
# Autogenerated by Thrift for thrift/compiler/test/fixtures/py3/src/module.thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#


from thrift.python.types import EnumMeta as __EnumMeta
import thrift.py3.types
import module.thrift_metadata

_fbthrift__module_name__ = "module.types"



class AnEnum(
    thrift.py3.types.CompiledEnum,
    metaclass=__EnumMeta,
):
    NOTSET = 0
    ONE = 1
    TWO = 2
    THREE = 3
    FOUR = 4

    __module__ = _fbthrift__module_name__
    __slots__ = ()

    @staticmethod
    def __get_metadata__():
        return module.thrift_metadata.gen_metadata_enum_AnEnum()

    @staticmethod
    def __get_thrift_name__():
        return "module.AnEnum"

    def _to_python(self):
        import importlib
        python_types = importlib.import_module(
            "module.thrift_types"
        )
        return python_types.AnEnum(self.value)

    def _to_py3(self):
        return self

    def _to_py_deprecated(self):
        return self.value




class AnEnumRenamed(
    thrift.py3.types.CompiledEnum,
    metaclass=__EnumMeta,
):
    name_ = 0
    value_ = 1
    renamed_ = 2

    __module__ = _fbthrift__module_name__
    __slots__ = ()

    @staticmethod
    def __get_metadata__():
        return module.thrift_metadata.gen_metadata_enum_AnEnumRenamed()

    @staticmethod
    def __get_thrift_name__():
        return "module.AnEnumRenamed"

    def _to_python(self):
        import importlib
        python_types = importlib.import_module(
            "module.thrift_types"
        )
        return python_types.AnEnumRenamed(self.value)

    def _to_py3(self):
        return self

    def _to_py_deprecated(self):
        return self.value




class Flags(
    thrift.py3.types.Flag,
    metaclass=__EnumMeta,
):
    flag_A = 1
    flag_B = 2
    flag_C = 4
    flag_D = 8

    __module__ = _fbthrift__module_name__
    __slots__ = ()

    @staticmethod
    def __get_metadata__():
        return module.thrift_metadata.gen_metadata_enum_Flags()

    @staticmethod
    def __get_thrift_name__():
        return "module.Flags"

    def _to_python(self):
        import importlib
        python_types = importlib.import_module(
            "module.thrift_types"
        )
        return python_types.Flags(self.value)

    def _to_py3(self):
        return self

    def _to_py_deprecated(self):
        return self.value





class __BinaryUnionType(
    thrift.py3.types.CompiledEnum,
    metaclass=__EnumMeta,
):
    iobuf_val = 1
    EMPTY = 0

    __module__ = _fbthrift__module_name__
    __slots__ = ()

